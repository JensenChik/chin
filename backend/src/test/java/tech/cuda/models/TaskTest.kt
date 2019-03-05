package tech.cuda.models

import org.junit.Test

import org.junit.Assert.*
import org.junit.Before
import tech.cuda.tools.DataMocker

/**
 * Created by Jensen on 19-1-26.
 */
class TaskTest {

    @Before
    fun init() {
        rebuildTables()
        DataMocker.loadMockedData()
    }

    @Test
    fun shouldScheduledToday() {
    }
}