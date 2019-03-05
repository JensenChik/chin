package tech.cuda.models.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.enums.InstanceStatus
import tech.cuda.models.mappers.Instance
import tech.cuda.models.mappers.Instances

/**
 * Created by Jensen on 19-3-5.
 */

object InstanceService {
    fun getInstances(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = Instances.select {
            Instances.removed.neq(true)
        }.orderBy(Instances.id to true).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }

    fun getInstancesByJobId(jobId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = Instances.select {
            Instances.removed.neq(true) and Instances.job.eq(jobId)
        }.orderBy(Instances.id to true).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }

    fun getInstancesByStatus(status: InstanceStatus, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = Instances.select {
            Instances.removed.neq(true) and Instances.status.eq(status)
        }.orderBy(Instances.id to true).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }
}