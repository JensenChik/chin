# coding=utf-8
import smtplib
from email.mime.text import MIMEText
from email.header import Header

sender = ''
receiver = ''
subject = '测试邮箱'
smtpserver = ''

msg = MIMEText('邮箱登陆成功：\n日志详情.....', 'plain', 'utf-8')  # 中文需参数‘utf-8'，单字节字符不需要
msg['Subject'] = Header(subject, 'utf-8')
msg['From'] = '调度器<>'
msg['To'] = receiver
smtp = smtplib.SMTP()
smtp.connect('', port=25)
smtp.login('', '')
smtp.sendmail(msg['From'], msg['To'].split(","), msg.as_string())
smtp.close()
