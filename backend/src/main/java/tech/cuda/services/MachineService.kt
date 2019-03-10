package tech.cuda.services

import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.exceptions.StringOutOfLengthException
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
        return when {
            name.length > MachineTable.NAME_MAX_LEN ->
                throw StringOutOfLengthException("name", MachineTable.NAME_MAX_LEN)
            ip.length > MachineTable.IP_MAX_LEN ->
                throw StringOutOfLengthException("ip", MachineTable.IP_MAX_LEN)
            mac.length > MachineTable.MAC_MAX_LEN ->
                throw StringOutOfLengthException("mac", MachineTable.MAC_MAX_LEN)
            else -> Machine.new {
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
}
