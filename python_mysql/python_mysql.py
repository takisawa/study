import mysql.connector

config = {
    'user': 'USER',
    'password': 'PASSWORD',
    'host': 'HOST',
    'database' : 'DATABASE',
}

connection = mysql.connector.connect(**config)
cursor = connection.cursor()
cursor.execute('select * from books')

for row in cursor.fetchall():
    print(row)

cursor.close()
connection.close()
