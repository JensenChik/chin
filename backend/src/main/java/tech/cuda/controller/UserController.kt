package tech.cuda.controller

import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.ResponseBody
import org.springframework.web.bind.annotation.RestController

/**
 * Created by Jensen on 18-6-15.
 * 用户操作controller
 */
@RestController
@EnableAutoConfiguration
@RequestMapping("/api/user")
class UserController {

    @RequestMapping("listing")
    @ResponseBody
    fun listing(): String {
        return "user"
    }

    @RequestMapping("login")
    @ResponseBody
    fun login(): String {
        return "登陆"
    }

    @RequestMapping("logout")
    @ResponseBody
    fun logout(): String {
        return "登出"
    }

    @RequestMapping("create")
    @ResponseBody
    fun create(): String {
        return "增加"
    }

    @RequestMapping("remove")
    @ResponseBody
    fun remove(): String {
        return "删除"
    }

    @RequestMapping("update")
    @ResponseBody
    fun update(): String {
        return "更新"
    }

}