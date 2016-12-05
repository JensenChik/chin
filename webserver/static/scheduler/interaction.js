/**
 * Created by suit on 16-12-5.
 */

function kill_task(task_id, version) {
    $.post('/kill_task', {'task_id': task_id, 'version': version}, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            location.reload();
        } else {
            alert('中止任务执行过程中出错');
        }
    });
}
