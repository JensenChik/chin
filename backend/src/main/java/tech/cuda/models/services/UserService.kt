package tech.cuda.models.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.models.mappers.User
import tech.cuda.models.mappers.Users

/**
 * Created by Jensen on 19-3-5.
 */

object UserService {
    fun getUsers(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = Users.select {
            Users.removed.neq(true)
        }.orderBy(Users.id to true).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }

    fun getUsersByGroupId(groupId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = Users.select {
            Users.removed.neq(true) and Users.group.eq(groupId)
        }.orderBy(Users.id to true).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }
}