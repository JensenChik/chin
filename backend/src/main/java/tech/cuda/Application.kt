package tech.cuda

import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter
import tech.cuda.models.rebuildTables
import java.util.*
import java.util.regex.Pattern
import javax.servlet.*
import javax.servlet.http.HttpServletRequest

/**
 * Created by Jensen on 18-6-15.
 */
@SpringBootApplication
open class Application : WebMvcConfigurerAdapter(), Filter {

    private val REWRITE_TO = "rewriteUrl" //需要rewrite到的目的地址
    private val REWRITE_PATTERNS = "urlPatterns" //拦截的url,url通配符之前用英文分号隔开
    private var urlPatterns: Set<String>? = null//配置url通配符
    private var rewriteTo: String? = null


    override fun init(cfg: FilterConfig) {
        val exceptUrl = cfg.getInitParameter(REWRITE_PATTERNS) ?: ""
        rewriteTo = cfg.getInitParameter(REWRITE_TO)
        urlPatterns = if (exceptUrl.isNotEmpty()) {
            Collections.unmodifiableSet(HashSet(Arrays.asList(
                    *exceptUrl.split(";".toRegex()).dropLastWhile({
                        it.isEmpty()
                    }).toTypedArray()
            )))
        } else {
            emptySet()
        }
    }

    override fun doFilter(req: ServletRequest, resp: ServletResponse, chain: FilterChain) {
        fun isMatches(patterns: Set<String>?, url: String): Boolean {
            if (null == patterns) {
                return false
            }
            for (str in patterns) {
                if (str.endsWith("/*")) {
                    val name = str.substring(0, str.length - 2)
                    if (url.contains(name)) {
                        return true
                    }
                } else {
                    val pattern = Pattern.compile(str)
                    if (pattern.matcher(url).matches()) {
                        return true
                    }
                }
            }
            return false
        }

        val request = req as HttpServletRequest
        val servletPath = request.servletPath
        val context = request.contextPath
        //匹配的路径重写
        if (isMatches(urlPatterns, servletPath)) {
            req.getRequestDispatcher("$context/$rewriteTo").forward(req, resp)
        } else {
            chain.doFilter(req, resp)
        }
    }

    override fun destroy() {}


    override fun addResourceHandlers(registry: ResourceHandlerRegistry?) {
        registry!!.addResourceHandler("/static/**").addResourceLocations("classpath:/static/")
        super.addResourceHandlers(registry)
    }
}

fun main(args: Array<String>) {
//    rebuildTables()
    SpringApplication.run(Application::class.java, *args)
}

