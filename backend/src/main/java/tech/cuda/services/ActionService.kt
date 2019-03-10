package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.exceptions.StringOutOfLengthException
import tech.cuda.models.Action
import tech.cuda.models.ActionTable
import tech.cuda.models.User

/**
 * Created by Jensen on 19-3-4.
 */
object ActionService {

    fun getOneById(id: Int): Action? {
        return Action.findById(id)
    }


    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Action> {
        val query = ActionTable.select {
            ActionTable.removed.neq(true)
        }.orderBy(ActionTable.id to false).limit(pageSize, offset = page * pageSize)
        return Action.wrapRows(query).toList()
    }

    fun getManyByUserId(userId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Action> {
        val query = ActionTable.select {
            ActionTable.user eq userId and ActionTable.removed.neq(true)
        }.orderBy(ActionTable.id to false).limit(pageSize, offset = page * pageSize)
        return Action.wrapRows(query).toList()
    }

    fun createOne(user: User, detail: String): Action {
        val now = DateTime.now()
        return if (detail.length > ActionTable.DETAIL_MAX_LEN)
            throw StringOutOfLengthException("detail", ActionTable.DETAIL_MAX_LEN)
        else Action.new {
            this.user = user
            this.detail = detail
            removed = false
            createTime = now
            updateTime = now
        }
    }

}