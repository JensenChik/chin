package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object GroupTable : IntIdTable() {
    override val tableName: String
        get() = "groups"

    val name = varchar(name = "name", length = 256).index()
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}

class Group(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Group>(GroupTable)

    var name by GroupTable.name
    var removed by GroupTable.removed
    var createTime by GroupTable.createTime
    var updateTime by GroupTable.updateTime
}