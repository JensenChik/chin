package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.models.User
import tech.cuda.models.Users

/**
 * Created by Jensen on 19-3-5.
 */

object UserService {
    fun getOneById(id: Int): User? {
        return User.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = Users.select {
            Users.removed.neq(true)
        }.orderBy(Users.id to true).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }

    fun getManyByGroupId(groupId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = Users.select {
            Users.removed.neq(true) and Users.group.eq(groupId)
        }.orderBy(Users.id to true).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }

    fun createOne() {

    }
}