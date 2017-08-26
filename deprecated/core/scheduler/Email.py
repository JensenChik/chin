# coding=utf-8
import smtplib
from email.mime.text import MIMEText
from email.header import Header
import ConfigParser


class Email:
    def __init__(self):
        cf = ConfigParser.ConfigParser()
        cf.read('chin.ini')
        self.host = cf.get('email', 'host')
        self.port = cf.get('email', 'port')
        self.sender = cf.get('email', 'sender')
        self.passwd = cf.get('email', 'passwd')
        self.receiver = cf.get('email', 'receiver')

    def send(self, subject, msg):
        msg = MIMEText(msg, 'plain', 'utf-8')
        msg['Subject'] = Header(subject, 'utf-8')
        msg['From'] = '调度器<{}>'.format(self.sender)
        msg['To'] = self.receiver
        server = smtplib.SMTP()
        server.connect(self.host, self.port)
        server.login(self.sender, self.passwd)
        server.sendmail(msg['From'], msg['To'].split(','), msg.as_string())
        server.close()
