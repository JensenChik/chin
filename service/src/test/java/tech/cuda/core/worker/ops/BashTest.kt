package tech.cuda.core.worker.ops

import org.apache.commons.lang3.RandomStringUtils
import org.junit.Assert.*
import org.junit.Test
import tech.cuda.models.Task

/**
 * Created by Jensen on 18-6-19.
 */
class BashTest {
    @Test
    fun testSuccess() {
        val bash = Bash("echo \$JAVA_HOME")
        bash.run()
        bash.wait()
        assertTrue(bash.isFinish)
        assertTrue(bash.isSuccess)
    }

    @Test
    fun testInitByTask(){
    }

    @Test
    fun testFailed() {
        val bash = Bash("java")
        bash.run()
        bash.wait()
        assertTrue(bash.isFinish)
        assertFalse(bash.isSuccess)
    }

    @Test
    fun testOutputCorrectly() {
        val bash = Bash("echo anything")
        bash.run()
        assertTrue(bash.wait())
        assertEquals("anything\n", bash.output)
    }

    @Test
    fun testRedirectCorrectly() {
        val filename = RandomStringUtils.randomAlphabetic(32) + ".log"

        var bash = Bash("echo anything > /tmp/$filename")
        bash.run()
        assertTrue(bash.wait())

        bash = Bash("ls /tmp/$filename | wc -l")
        bash.run()
        assertTrue(bash.wait())
        assertEquals(1, Integer.parseInt(bash.output.trim()))

        bash = Bash("rm /tmp/$filename")
        bash.run()
        assertTrue(bash.wait())

        bash = Bash("ls /tmp -al | grep $filename | wc -l")
        bash.run()
        assertTrue(bash.wait())
        assertEquals(0, Integer.parseInt(bash.output.trim()))
    }

    @Test
    @Throws(InterruptedException::class)
    fun testRunningAndKilling() {
        val bash = Bash("echo anything && sleep 100 && sleep 20")
        bash.run()
        Thread.sleep(1000)
        assertFalse(bash.isFinish)
        assertEquals("anything\n", bash.output)

        bash.kill()
        Thread.sleep(1000)
        assertTrue(bash.isFinish)
        assertFalse(bash.isSuccess)
    }

}