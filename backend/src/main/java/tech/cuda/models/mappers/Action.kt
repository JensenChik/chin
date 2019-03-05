package tech.cuda.models.mappers

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Actions : IntIdTable() {
    override val tableName: String
        get() = "actions"

    val user = reference(name = "user_id", foreign = Users)
    val type = varchar(name = "type", length = 16)
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}

class Action(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Action>(Actions)

    var user by User referencedOn Actions.user
    var type by Actions.type
    var removed by Actions.removed
    var createTime by Actions.createTime
    var updateTime by Actions.updateTime
}