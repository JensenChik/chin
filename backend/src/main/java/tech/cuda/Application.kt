package tech.cuda

import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.web.servlet.FilterRegistrationBean
import org.springframework.context.annotation.Bean
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter
import javax.servlet.*
import javax.servlet.http.HttpServletRequest

/**
 * Created by Jensen on 18-6-15.
 */


@SpringBootApplication
open class Application : WebMvcConfigurerAdapter(), Filter {

    override fun init(cfg: FilterConfig) {}

    override fun doFilter(req: ServletRequest, resp: ServletResponse, chain: FilterChain) {
        req as HttpServletRequest
        val isBackendRouter = req.servletPath.startsWith("/dist/")
                || req.servletPath in listOf("/index.html", "/favicon.ico")
                || req.servletPath.startsWith("/api/")
        if (isBackendRouter) {
            chain.doFilter(req, resp)
        } else {
            req.getRequestDispatcher("${req.contextPath}/index.html").forward(req, resp)
        }
    }

    override fun destroy() {}


    override fun addResourceHandlers(registry: ResourceHandlerRegistry?) {
        registry!!.addResourceHandler("/static/**").addResourceLocations("classpath:/static/")
        super.addResourceHandlers(registry)
    }


    @Bean
    open fun registerFilter(): FilterRegistrationBean {
        val registration = FilterRegistrationBean()
        registration.setName("RewriteFilter")
        registration.filter = Application()
        registration.addUrlPatterns("/*")
        registration.order = 1
        return registration
    }

}

fun main(args: Array<String>) {
    // todo: 添加命令行参数
//    rebuildTables()
    SpringApplication.run(Application::class.java, *args)
}

