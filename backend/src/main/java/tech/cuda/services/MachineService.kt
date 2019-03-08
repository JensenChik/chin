package tech.cuda.services

import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.models.Machine
import tech.cuda.models.MachineTable

/**
 * Created by Jensen on 19-3-5.
 */

object MachineService {
    fun getOneById(id: Int): Machine? {
        return Machine.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Machine> {
        val query = MachineTable.select {
            MachineTable.removed.neq(true)
        }.orderBy(MachineTable.id to false).limit(pageSize, offset = page * pageSize)
        return Machine.wrapRows(query).toList()

    }

    fun createOne(name: String, ip: String, mac: String): Machine {
        val now = DateTime.now()
        return Machine.new {
            this.name = name
            this.ip = ip
            this.mac = mac
            cpuLoad = -1
            memLoad = -1
            diskLoad = -1
            removed = false
            createTime = now
            updateTime = now
        }
    }
}