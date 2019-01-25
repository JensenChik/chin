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
@RequestMapping("/api/task")
class TaskController {

    @RequestMapping("listing")
    @ResponseBody
    fun listing(): String {
        return "task"
    }

    @RequestMapping("add")
    @ResponseBody
    fun add(): String {
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