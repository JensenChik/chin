package tech.cuda.core.worker

import org.joda.time.DateTime
import tech.cuda.models.InstanceStatus
import tech.cuda.ops.Bash
import tech.cuda.services.InstanceService

/**
 * Created by Jensen on 18-10-25.
 * 只能写 instance 这个表
 */
object InstanceTracker {

    fun startOperatorForWaitingInstance() {

    }


    fun stopOperatorForKillingInstance() {

    }

    fun checkOperatorStatusForRunningInstance() {

    }

    fun closeForFinishedInstance() {

    }


    fun serve() {
        val bashPool = mutableMapOf<Int, Bash>()
        while (true) {
            InstanceService.getManyByStatus(InstanceStatus.Waiting).forEach {
                val bash = InstanceService.start(it)
                bashPool[it.id.value] = bash
            }
            Thread.sleep(1000)
            bashPool.filter {
                it.value.isFinish
            }.forEach {
                val instance = InstanceService.getOneById(it.key)!!
                instance.status = if (it.value.isSuccess) InstanceStatus.Success else InstanceStatus.Failed
                instance.output = it.value.output
                instance.updateTime = DateTime.now()
            }
            Thread.sleep(1000)

        }

    }

}