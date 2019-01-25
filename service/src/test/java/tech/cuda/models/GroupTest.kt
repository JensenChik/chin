package tech.cuda.models

import org.jetbrains.exposed.sql.transactions.transaction
import org.joda.time.DateTime
import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.enums.ScheduleType

/**
 * Created by Jensen on 19-1-26.
 */
class GroupTest {


    @Before
    fun init() {
        rebuildTables()
        transaction {
            val group1 = Group.new {
                name = "权限组1"
                createTime = DateTime.now()
                updateTime = DateTime.now()
            }

            val group2 = Group.new {
                name = "权限组2"
                createTime = DateTime.now()
                updateTime = DateTime.now()
            }

            for (i in 1..10) {
                User.new {
                    group = group1
                    name = "权限组1,用户${i}"
                    password = "xxxx"
                    email = "xxx"
                    createTime = DateTime.now()
                    updateTime = DateTime.now()
                }
            }
            for (i in 1..5) {
                User.new {
                    group = group2
                    name = "权限组2,用户${i}"
                    password = "yyyy"
                    email = "yyy"
                    createTime = DateTime.now()
                    updateTime = DateTime.now()
                }
            }
        }

    }


    @Test
    fun getUsers() {
        transaction {
            for (i in 1..10) {
                val users = User.findById(i)!!.group.inclusiveUsers
                assertTrue(users.count() == 10)
                assertTrue(users.sumBy { it.id.value } == 55)
            }
            for (i in 11..15) {
                val users = User.findById(i)!!.group.inclusiveUsers
                assertTrue(users.count() == 5)
                assertTrue(users.sumBy { it.id.value } == 65)
            }

        }

    }

}