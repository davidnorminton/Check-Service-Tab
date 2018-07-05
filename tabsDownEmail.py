#/bin/env python

import sys
import smtplib
from email.MIMEMultipart import MIMEMultipart
from email.MIMEText import MIMEText
 
 
fromaddr = ""
toaddr = ""
msg = MIMEMultipart('alternative')
msg['From'] = fromaddr
msg['To'] = toaddr
msg['Subject'] = "cs-catering Service tab is down"
 
body = "<html><head></head><body><h1>Services tab is down</h1>" + sys.argv[1] + "</body></html>"
msg.attach(MIMEText(body, 'html'))
 
server = smtplib.SMTP('smtp.gmail.com', 587)
server.starttls()
server.login(fromaddr, SMTP_PASSWORD_GOES_HERE)
text = msg.as_string()
server.sendmail(fromaddr, toaddr, text)
server.quit()
print(body)


