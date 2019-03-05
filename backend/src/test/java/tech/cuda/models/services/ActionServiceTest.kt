package tech.cuda.models.services

import org.jetbrains.exposed.sql.transactions.transaction
import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.models.rebuildTables
import tech.cuda.tools.DataMocker

/**
 * Created by Jensen on 19-3-4.
 */
class ActionServiceTest {

    @Before
    fun init() {
        rebuildTables()
        DataMocker.load(listOf("groups", "users", "actions"))
    }

    @Test
    fun getActionsByUserId() {
        transaction {
            assertEquals("总数不等", 4, ActionService.getActionsByUserId(1).size)
            assertEquals("总数不等", 2, ActionService.getActionsByUserId(2).size)
            assertEquals("总数不等", 1, ActionService.getActionsByUserId(3).size)
            assertEquals("分页总数不等", 3, ActionService.getActionsByUserId(1, page = 0, pageSize = 3).size)
            assertEquals("分页总数不等", 1, ActionService.getActionsByUserId(1, page = 1, pageSize = 3).size)
        }
    }
}