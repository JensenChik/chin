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

function run_at_once(task_id) {
    $.post('/run_at_once', {'task_id': task_id}, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            location.href = result.data.url;
        }
    });
}

function reverse_task_valid(task_id, valid) {
    $.post('/reverse_task_valid', {'task_id': task_id, 'valid': valid}, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            location.href = result.data.url;
        }
    });
}

function get_user_detail(user_name) {
    $.post('/get_user_detail', {'user_name': user_name}, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            $('#email').val(result.data.email);
        } else if (result.status == 'failed') {
            alert(result.err_info);
        }
    });
}

function modify_user(user_name, old_password, new_password, email) {
    var user = {
        "user_name": user_name,
        "old_password": old_password,
        "new_password": new_password,
        "email": email
    };
    $.post('/modify_user', user, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            location.href = result.data.href;
        } else if (result.status == 'failed') {
            alert(result.err_info);
        }
    });
}
