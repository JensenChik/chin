package tech.cuda.services

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
    fun setUp() {
        rebuildTables()
        DataMocker.load("groups", "users", "actions")
    }

    @Test
    fun getActionsByUserId() {
        transaction {
            assertEquals("总数不等", 4, ActionService.getManyByUserId(1).size)
            assertEquals("总数不等", 2, ActionService.getManyByUserId(2).size)
            assertEquals("总数不等", 1, ActionService.getManyByUserId(3).size)
            assertEquals("分页总数不等", 3, ActionService.getManyByUserId(1, page = 0, pageSize = 3).size)
            assertEquals("分页总数不等", 1, ActionService.getManyByUserId(1, page = 1, pageSize = 3).size)
        }
    }

    @Test
    fun createOne(){
        transaction {
            val user = UserService.getOneById(1)!!
            val action = ActionService.createOne(user, "增加一个任务")
            val actionToBeChecked = ActionService.getOneById(action.id.value)!!
            assertEquals("action 没有被正确插入", "增加一个任务", actionToBeChecked.detail)
            assertEquals("action 没有被正确插入", user.id, actionToBeChecked.user.id)
        }
    }
}