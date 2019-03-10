package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.exceptions.StringOutOfLengthException
import tech.cuda.models.Group
import tech.cuda.models.User
import tech.cuda.models.UserTable

/**
 * Created by Jensen on 19-3-5.
 */

object UserService {
    fun getOneById(id: Int): User? {
        return User.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = UserTable.select {
            UserTable.removed.neq(true)
        }.orderBy(UserTable.id to false).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }

    fun getManyByGroupId(groupId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<User> {
        val query = UserTable.select {
            UserTable.removed.neq(true) and UserTable.group.eq(groupId)
        }.orderBy(UserTable.id to false).limit(pageSize, offset = page * pageSize)
        return User.wrapRows(query).toList()
    }

    fun createOne(group: Group, name: String, password: String, email: String): User {
        val now = DateTime.now()
        return when {
            name.length > UserTable.NAME_MAX_LEN ->
                throw StringOutOfLengthException("name", UserTable.NAME_MAX_LEN)
            password.length > UserTable.PASSWORD_MAX_LEN ->
                throw StringOutOfLengthException("password", UserTable.NAME_MAX_LEN)
            email.length > UserTable.EMAIL_MAX_LEN ->
                throw StringOutOfLengthException("email", UserTable.EMAIL_MAX_LEN)
            else -> User.new {
                this.group = group
                this.name = name
                this.password = password
                this.email = email
                removed = false
                createTime = now
                updateTime = now
            }
        }


    }
}