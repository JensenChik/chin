package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object UserTable : IntIdTable() {
    override val tableName: String
        get() = "users"

    val group = reference(name = "group_id", foreign = GroupTable)
    val name = varchar(name = "name", length = 256).index()
    val password = varchar(name = "password", length = 256)
    val email = varchar(name = "email", length = 256)
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class User(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<User>(UserTable)
    var group by Group referencedOn UserTable.group
    var name by UserTable.name
    var password by UserTable.password
    var email by UserTable.email
    var removed by UserTable.removed
    var createTime by UserTable.createTime
    var updateTime by UserTable.updateTime
}
