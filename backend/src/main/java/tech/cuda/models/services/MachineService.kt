package tech.cuda.models.services

import org.jetbrains.exposed.sql.select
import tech.cuda.models.mappers.Machine
import tech.cuda.models.mappers.Machines

/**
 * Created by Jensen on 19-3-5.
 */

object MachineService {
    fun getMachines(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Machine> {
        val query = Machines.select {
            Machines.removed.neq(true)
        }.orderBy(Machines.id to true).limit(pageSize, offset = page * pageSize)
        return Machine.wrapRows(query).toList()

    }
}