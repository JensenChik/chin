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

function fill_log_to_div(task_id, version, div_id) {
    document.getElementById(div_id).innerText = '';
    $.post('/get_log_detail', {'task_id': task_id, 'version': version}, function (result) {
        document.getElementById(div_id).innerText = result;
    });
}

function rerun(task_id, version) {
    $.post('/rerun', {'task_id': task_id, 'version': version}, function () {
        location.reload();
    });
}

function list_instance_log(task_id) {
    location.href = '/list_instance_log?task_id=' + task_id;
}
