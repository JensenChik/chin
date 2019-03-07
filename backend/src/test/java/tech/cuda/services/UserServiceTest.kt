package tech.cuda.services

import org.jetbrains.exposed.sql.transactions.transaction
import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.models.User
import tech.cuda.models.rebuildTables
import tech.cuda.tools.DataMocker

/**
 * Created by Jensen on 19-3-6.
 */
class UserServiceTest {

    @Before
    fun setUp() {
        rebuildTables()
        DataMocker.load(listOf("groups", "users"))
    }

    @Test
    fun getUsers() {
    }

//    @Test
//    fun getUsersByGroupId() {
//        transaction {
//            for (i in 1..10) {
//                UserService.getOneById(i).group
//                val users = User.findById(i)!!.group.inclusiveUsers
//                assertTrue(users.count() == 10)
//                assertTrue(users.sumBy { it.id.value } == 55)
//            }
//            for (i in 11..15) {
//                val users = User.findById(i)!!.group.inclusiveUsers
//                assertTrue(users.count() == 5)
//                assertTrue(users.sumBy { it.id.value } == 65)
//            }
//
//        }
//    }
}