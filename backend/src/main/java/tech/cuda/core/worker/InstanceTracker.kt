package tech.cuda.core.worker

import tech.cuda.models.InstanceStatus
import tech.cuda.ops.Bash
import tech.cuda.services.InstanceService

/**
 * Created by Jensen on 18-10-25.
 * 只能写 instance 这个表
 */
object InstanceTracker {

    fun serve() {
        val bashPool = mutableMapOf<Int, Bash>()
        while (true) {
            // start for waiting instance
            InstanceService.getManyByStatus(InstanceStatus.Waiting).forEach {
                val bash = InstanceService.start(it)
                bashPool[it.id.value] = bash
            }
            Thread.sleep(1000)

            // stop for killing instance
            InstanceService.getManyByStatus(InstanceStatus.Killing).forEach {
                bashPool[it.id.value]?.kill()
            }
            Thread.sleep(1000)

            // close for finish bash
            bashPool.filter {
                it.value.isFinish
            }.forEach {
                val status = if (it.value.isSuccess) InstanceStatus.Success else InstanceStatus.Failed
                InstanceService.updateById(it.key, status, it.value.output)
            }
            Thread.sleep(1000)

        }

    }

}