package tech.cuda.models.services

import tech.cuda.models.mappers.Action
import tech.cuda.models.mappers.Actions

/**
 * Created by Jensen on 19-3-4.
 */
object ActionService {
    fun getActionsByUserId(id: Int, page: Int = 0, pageSize: Int = 10): List<Action> {
        return Action.find {
            Actions.user eq id
        }.sortedBy { it.id }.subList(page * pageSize, (page + 1) * pageSize)
    }

    fun getActionByPage(page: Int, pageSize: Int) {

    }
}