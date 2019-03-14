package tech.cuda.services

import com.mysql.jdbc.Blob
import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.models.Instance
import tech.cuda.models.InstanceStatus
import tech.cuda.models.InstanceTable
import tech.cuda.models.Job
import tech.cuda.ops.Bash
import javax.sql.rowset.serial.SerialBlob

/**
 * Created by Jensen on 19-3-5.
 */

object InstanceService {

    fun getOneById(id: Int): Instance? {
        return Instance.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = InstanceTable.select {
            InstanceTable.removed.neq(true)
        }.orderBy(InstanceTable.id to false).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }

    fun getManyByJobId(jobId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = InstanceTable.select {
            InstanceTable.removed.neq(true) and InstanceTable.job.eq(jobId)
        }.orderBy(InstanceTable.id to false).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }

    fun getManyByStatus(status: InstanceStatus, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Instance> {
        val query = InstanceTable.select {
            InstanceTable.removed.neq(true) and InstanceTable.status.eq(status)
        }.orderBy(InstanceTable.id to false).limit(pageSize, offset = page * pageSize)
        return Instance.wrapRows(query).toList()
    }

    fun createOneForJob(job: Job): Instance {
        val now = DateTime.now()
        return Instance.new {
            this.job = job
            status = InstanceStatus.Waiting
            removed = false
            createTime = now
            updateTime = now
        }
    }

    fun updateById(id: Int, status: InstanceStatus? = null, output: String? = null): Instance {
        val instance = this.getOneById(id)!!
        val now = DateTime.now()
        if (status != null) {
            instance.status = status
            instance.updateTime = now
        }
        if (output != null) {
            instance.output = SerialBlob(output.toByteArray())
            instance.updateTime = now
        }
        return instance
    }
    

    fun removeById(id: Int): Instance {
        val instance = this.getOneById(id)!!
        instance.removed = true
        instance.updateTime = DateTime.now()
        return instance

    }

    fun start(instance: Instance): Bash {
        instance.status = InstanceStatus.Running
        val bash = Bash(instance.job.task.command)
        bash.run()
        return bash
    }

    fun kill(instance: Instance) {

    }
}