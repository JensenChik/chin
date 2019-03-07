package tech.cuda.services

import org.jetbrains.exposed.sql.select
import tech.cuda.models.Group
import tech.cuda.models.Groups

/**
 * Created by Jensen on 19-3-5.
 */
object GroupService {

    fun getOneById(id: Int): Group? {
        return Group.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Group> {
        val query = Groups.select {
            Groups.removed.neq(true)
        }.orderBy(Groups.id to true).limit(pageSize, offset = page * pageSize)
        return Group.wrapRows(query).toList()
    }

    fun createOne() {

    }
}