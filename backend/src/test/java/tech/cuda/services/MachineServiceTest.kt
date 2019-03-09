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
class MachineServiceTest {

    @Before
    fun setUp() {
        rebuildTables()
        DataMocker.load("machines")
    }

    @Test
    fun `query multi machines`() {
        transaction {
            assertEquals("机器总数不等", 6, MachineService.getMany().size)
        }
    }
}