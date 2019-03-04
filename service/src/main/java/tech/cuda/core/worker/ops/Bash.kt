package tech.cuda.core.worker.ops

import org.apache.commons.exec.*
import org.apache.commons.exec.environment.EnvironmentUtils
import org.apache.commons.io.IOUtils
import tech.cuda.models.mappers.Task
import java.io.ByteArrayOutputStream
import java.io.IOException

/**
 * Created by Jensen on 18-6-15.
 * bash 命令调度
 */
const val EXIT_VALUE = 0

class Bash : Operator {
    private val commandLine: CommandLine = CommandLine("bash")
    private val std: ByteArrayOutputStream
    private val handle: DefaultExecuteResultHandler
    private val killer: ExecuteWatchdog
    private val executor: DefaultExecutor

    val output: String
        get() {
            return this.std.toString()
        }

    val isFinish: Boolean
        get() {
            return this.handle.hasResult() || this.killer.killedProcess()
        }

    val isSuccess: Boolean
        get() {
            return try {
                this.isFinish && this.handle.exitValue == EXIT_VALUE
            } catch (e: IllegalStateException) {
                false
            }
        }


    constructor(command: String) {
        commandLine.addArgument("-c")
        commandLine.addArgument(command, false)

        this.std = ByteArrayOutputStream()
        this.handle = DefaultExecuteResultHandler()
        this.killer = ExecuteWatchdog(Long.MAX_VALUE)

        this.executor = DefaultExecutor()
        this.executor.setExitValue(EXIT_VALUE)
        this.executor.streamHandler = PumpStreamHandler(std, std)
        this.executor.watchdog = killer
    }

    constructor(task: Task) : this(task.command)


    fun run() {
        try {
            executor.execute(commandLine, EnvironmentUtils.getProcEnvironment(), this.handle)
        } catch (e: IOException) {
            this.executor.setExitValue(-1)
            // todo: 日志打印
        }
    }

    fun wait(): Boolean {
        this.handle.waitFor()
        return this.isSuccess
    }

    fun kill() {
        try {
            // 先获取看门狗守护的 process
            var field = ExecuteWatchdog::class.java.getDeclaredField("process")
            field.isAccessible = true
            val process = field.get(this.killer)

            // 然后获取 process 的 pid
            field = process.javaClass.getDeclaredField("pid")
            field.isAccessible = true
            val pid = field.getInt(process)

            // 再获取 pid 所关联的进程树, 逐个kill掉
            val pstree = Runtime.getRuntime().exec("pstree -p $pid")
            pstree.waitFor()

            """(?<=\()[^)]+""".toRegex().findAll(IOUtils.toString(pstree.inputStream)).forEach {
                val kill = Runtime.getRuntime().exec("kill -9 ${it.value}")
                kill.waitFor()
            }


        } catch (e: IllegalAccessException) {
            e.printStackTrace()
            // todo 日志打印
        } catch (e: InterruptedException) {
            e.printStackTrace()
        } catch (e: NoSuchFieldException) {
            e.printStackTrace()
        } catch (e: IOException) {
            e.printStackTrace()
        }

    }


}