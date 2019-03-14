package tech.cuda.controller

import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.ResponseBody
import org.springframework.web.bind.annotation.RestController

/**
 * Created by Jensen on 18-6-15.
 */

@RestController
@EnableAutoConfiguration
@RequestMapping("/")
class IndexController {

    @RequestMapping("home")
    @ResponseBody
    fun home(): String {
        return "index"
    }

}