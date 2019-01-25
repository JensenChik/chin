package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Users : IntIdTable() {
    override val tableName: String
        get() = "users"

    val group = reference(name = "group_id", foreign = Groups)
    val name = varchar(name = "name", length = 256).index()
    val password = varchar(name = "password", length = 256)
    val email = varchar(name = "email", length = 256)
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class User(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<User>(Users)

    var group by Group referencedOn Users.group
    var name by Users.name
    var password by Users.password
    var email by Users.email
    var removed by Users.removed
    var createTime by Users.createTime
    var updateTime by Users.updateTime

}
