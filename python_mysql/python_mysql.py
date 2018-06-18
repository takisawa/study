import mysql.connector

config = {
    'user': 'USER',
    'password': 'PASSWORD',
    'host': 'HOST',
    'database' : 'DATABASE',
}

connection = mysql.connector.connect(**config)


print("--- select all ---")
cursor = connection.cursor()
cursor.execute('select * from books')

for row in cursor.fetchall():
    print(row)

cursor.close()

print("--- select one ---")
cursor = connection.cursor()
cursor.execute('select * from books where id=%s', (2,))

row = cursor.fetchone()
print(row)

cursor.close()

connection.close()
