package tech.cuda.controller

import org.springframework.boot.autoconfigure.EnableAutoConfiguration
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.ResponseBody
import org.springframework.web.bind.annotation.RestController
import tech.cuda.models.ScheduleFormat
import tech.cuda.models.ScheduleType
import tech.cuda.services.TaskService
import tech.cuda.services.UserService

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

    @RequestMapping("create")
    @ResponseBody
    fun create(): String {
        val user = UserService.getOneById(1)!!
        val name = "创建任务"
        val task = TaskService.createOne(
                user = user,
                name = name,
                scheduleType = ScheduleType.Day,
                scheduleFormat = ScheduleFormat(hour = 0, minute = 10),
                command = "echo hello world"
        )
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