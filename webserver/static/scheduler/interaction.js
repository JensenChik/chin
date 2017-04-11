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

function join_queue(task_id) {
    $.post('/join_queue', {'task_id': task_id}, function (result) {
        result = JSON.parse(result);
        alert(result.info);
    })
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

function change_scheduled_format(scheduled_type) {
    $('#scheduled_type').html(scheduled_type);
    $('#scheduled_format').children().each(function () {
        $(this).hide()
    });
    $('#' + scheduled_type).show();
}

function load_task_detail(task_id) {
    $.post('/get_task_detail', {'task_id': task_id}, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            var task = result.data;
            $('#task_name').val(task.name);
            $('#command').val(task.command);
            $('#valid').val(task.valid.toString());
            $('#priority').val(task.priority);
            $('#rerun').val(task.rerun.toString());
            $('#rerun_times').val(task.rerun_times);
            $('#father_task').val(task.father_task.join('\n'));
            $('#machine_pool').val(task.machine_pool.join('\n'));
            if (task.scheduled_type == 'day') {
                change_scheduled_format('每日');
                $('#每日').children('.scheduled_hour').children().val(task.hour);
                $('#每日').children('.scheduled_minute').children().val(task.minute);
            } else if (task.scheduled_type == 'week') {
                change_scheduled_format('每周');
                $('#每周').children('.scheduled_weekday').children().val(task.weekday);
                $('#每周').children('.scheduled_hour').children().val(task.hour);
                $('#每周').children('.scheduled_minute').children().val(task.minute);
            } else if (task.scheduled_type == 'month') {
                change_scheduled_format('每月');
                $('#每月').children('.scheduled_day').children().val(task.day);
                $('#每月').children('.scheduled_hour').children().val(task.hour);
                $('#每月').children('.scheduled_minute').children().val(task.minute);
            } else if (task.scheduled_type == 'month') {
            } else if (task.scheduled_type == 'once') {
                change_scheduled_format('一次');
                $('#一次').children('.scheduled_year').children().val(task.year);
                $('#一次').children('.scheduled_month').children().val(task.month);
                $('#一次').children('.scheduled_day').children().val(task.day);
                $('#一次').children('.scheduled_hour').children().val(task.hour);
                $('#一次').children('.scheduled_minute').children().val(task.minute);
            }
        } else if (result.status == 'failed') {
            alert(result.err_info);
        }
    });
}


function submit_task(url, has_next) {
    var scheduled_type = $('#scheduled_type').html();
    var scheduled_time = $('#' + scheduled_type);
    scheduled_type = {'每周': 'week', '每日': 'day', '每月': 'month', '一次': 'once'}[scheduled_type];
    var task_meta = {
        'task_id': $('#task_id').val(),
        'task_name': $('#task_name').val(),
        'command': $('#command').val(),
        'valid': $('#valid').val(),
        'priority': $('#priority').val(),
        'rerun': $('#rerun').val(),
        'rerun_times': $('#rerun_times').val(),
        'machine_pool': $('#machine_pool').val(),
        'father_task': $('#father_task').val(),
        'scheduled_type': scheduled_type,
        'year': scheduled_time.children('.scheduled_year').children().val(),
        'month': scheduled_time.children('.scheduled_month').children().val(),
        'weekday': scheduled_time.children('.scheduled_weekday').children().val(),
        'day': scheduled_time.children('.scheduled_day').children().val(),
        'hour': scheduled_time.children('.scheduled_hour').children().val(),
        'minute': scheduled_time.children('.scheduled_minute').children().val(),
        'has_next': has_next
    };
    $.post(url, task_meta, function (result) {
        result = JSON.parse(result);
        if (result.status == 'success') {
            location.href = result.data.url;
        }
    });
}
