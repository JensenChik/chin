package tech.cuda.services

import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.models.Group
import tech.cuda.models.GroupTable

/**
 * Created by Jensen on 19-3-5.
 */
object GroupService {

    fun getOneById(id: Int): Group? {
        return Group.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Group> {
        val query = GroupTable.select {
            GroupTable.removed.neq(true)
        }.orderBy(GroupTable.id to false).limit(pageSize, offset = page * pageSize)
        return Group.wrapRows(query).toList()
    }

    fun createOne(name: String): Group {
        val now = DateTime.now()
        return Group.new {
            this.name = name
            this.removed = false
            this.createTime = now
            this.updateTime = now
        }
    }
}