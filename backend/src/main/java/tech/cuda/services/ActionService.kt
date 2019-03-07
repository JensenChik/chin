package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.models.Action
import tech.cuda.models.Actions

/**
 * Created by Jensen on 19-3-4.
 */
object ActionService {

    /**
     * @author Jensen
     * @param id: actions 表的主键 ID
     * @return action
     */
    fun getOneById(id: Int): Action? {
        return Action.findById(id)
    }


    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Action> {
        val query = Actions.select {
            Actions.removed.neq(true)
        }.orderBy(Actions.id to true).limit(pageSize, offset = page * pageSize)
        return Action.wrapRows(query).toList()
    }

    fun getManyByUserId(userId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Action> {
        val query = Actions.select {
            Actions.user eq userId and Actions.removed.neq(true)
        }.orderBy(Actions.id to true).limit(pageSize, offset = page * pageSize)
        return Action.wrapRows(query).toList()
    }

    fun createOne() {

    }

}