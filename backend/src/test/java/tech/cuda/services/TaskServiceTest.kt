package tech.cuda.services

import org.jetbrains.exposed.sql.transactions.transaction
import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.models.rebuildTables
import tech.cuda.tools.DataMocker

/**
 * Created by Jensen on 19-3-6.
 */
class TaskServiceTest {

    @Before
    fun setUp() {
        rebuildTables()
        DataMocker.load("groups", "users", "tasks")
    }

    @Test
    fun getTasks() {
        transaction {
            assertEquals("总记录数不等", 9, TaskService.getMany().size)
            assertEquals("分页记录数不等", 5, TaskService.getMany(pageSize = 5).size)
            assertEquals("分页记录数不等", 4, TaskService.getMany(page = 1, pageSize = 5).size)
        }
    }

    @Test
    fun getTasksByUserId() {
    }

    @Test
    fun getTasksByScheduleType() {
    }
}