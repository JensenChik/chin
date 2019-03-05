package tech.cuda.models

import org.jetbrains.exposed.sql.transactions.transaction
import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.models.mappers.User
import tech.cuda.tools.DataMocker

/**
 * Created by Jensen on 19-1-26.
 */
class GroupTest {


    @Before
    fun init() {
        rebuildTables()
        DataMocker.loadMockedData()
    }


    @Test
    fun getInclusiveUsers() {
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