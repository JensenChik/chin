package tech.cuda.controller

import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.ResponseBody

/**
 * Created by Jensen on 18-6-15.
 */

@EnableAutoConfiguration
class IndexController {

    @RequestMapping("/")
    @ResponseBody
    fun home(): String {
        return "index"
    }

}