package tech.cuda.models.services

import org.jetbrains.exposed.sql.select
import tech.cuda.models.mappers.Group
import tech.cuda.models.mappers.Groups

/**
 * Created by Jensen on 19-3-5.
 */
object GroupService {
    fun getGroups(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Group> {
        val query = Groups.select {
            Groups.removed.neq(true)
        }.orderBy(Groups.id to true).limit(pageSize, offset = page * pageSize)
        return Group.wrapRows(query).toList()
    }
}