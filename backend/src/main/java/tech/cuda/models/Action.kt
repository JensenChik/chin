package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object ActionTable : IntIdTable() {
    override val tableName: String
        get() = "actions"

    val user = reference(name = "user_id", foreign = UserTable)
    const val detailLength = 256
    val detail = varchar(name = "detail", length = detailLength)
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}

class Action(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Action>(ActionTable)

    var user by User referencedOn ActionTable.user
    var detail by ActionTable.detail
    var removed by ActionTable.removed
    var createTime by ActionTable.createTime
    var updateTime by ActionTable.updateTime
}