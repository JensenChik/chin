package tech.cuda.models.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.jetbrains.exposed.sql.selectAll
import tech.cuda.models.mappers.Action
import tech.cuda.models.mappers.Actions

/**
 * Created by Jensen on 19-3-4.
 */
object ActionService {
    fun getActionsByUserId(id: Int, page: Int = 0, pageSize: Int = 10): List<Action> {
        val query = Actions.select {
            Actions.user eq id and Actions.removed.neq(true)
        }.orderBy(Actions.id to true).limit(pageSize, offset = page * pageSize)
        return Action.wrapRows(query).toList()
    }

}