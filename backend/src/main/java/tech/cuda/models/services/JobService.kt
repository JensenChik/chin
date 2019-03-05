package tech.cuda.models.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.models.mappers.Job
import tech.cuda.models.mappers.Jobs

/**
 * Created by Jensen on 19-3-5.
 */

object JobService {
    fun getJobs(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Job> {
        val query = Jobs.select {
            Jobs.removed.neq(true)
        }.orderBy(Jobs.id to true).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }

    fun getJobsByTaskId(taskId: Int, page: Int, pageSize: Int): List<Job> {
        val query = Jobs.select {
            Jobs.removed.neq(true) and Jobs.task.eq(taskId)
        }.orderBy(Jobs.id to true).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }
}