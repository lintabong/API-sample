import sqlite3
from flask_restful import abort


def get_user():
    conn = sqlite3.connect('mydatabase.db')
    cursor = conn.execute("SELECT * FROM user")

    userlist = {}
    for row in cursor:
        userlist['user' + str(row[0])] = {
            'id': row[0],
            'username': row[1],
            'password': row[2],
            'level': row[3]
        }

    conn.close()

    return userlist


def abort_user(user_id):
    if user_id not in get_user():
        abort(404, message='user {} doesnt exist'.format(user_id))


def get_post():
    conn = sqlite3.connect('mydatabase.db')
    cursor = conn.execute("SELECT * FROM post")

    postlist = {}
    for row in cursor:
        postlist['post' + str(row[0])] = {
            'id': row[0],
            'author': row[1],
            'title': row[2],
            'createdAt': row[3],
            'content': row[4]
        }

    conn.close()

    return postlist


def abort_post(post_id):
    if post_id not in get_post():
        abort(404, message='post {} doesnt exist'.format(post_id))


def get_category():
    conn = sqlite3.connect('mydatabase.db')
    cursor = conn.execute("SELECT * FROM category")

    categorylist = {}
    for row in cursor:
        categorylist['category' + str(row[0])] = {
            'id': row[0],
            'title': row[1]
        }

    conn.close()

    return categorylist


def abort_category(category_id):
    if category_id not in get_category():
        abort(404, message='post {} doesnt exist'.format(category_id))


def get_post_category():
    conn = sqlite3.connect('mydatabase.db')
    cursor = conn.execute("SELECT * FROM post_category")

    pclist = {}
    for row in cursor:
        pclist['post_category' + str(row[0]) + str(row[1])] = {
            'post_id': row[0],
            'category_id': row[1]
        }

    conn.close()

    return pclist
