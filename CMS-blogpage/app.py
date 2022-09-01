from flask import Flask, request
from flask_restful import Resource, Api, abort
import sqlite3
from datetime import datetime

app = Flask(__name__)
api = Api(app)


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
        abort(404, message='todo {} doesnt exist'.format(user_id))


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


class UserAdd(Resource):
    def get(self):
        return get_user()

    def post(self):
        username = request.args.get('username')
        password = request.args.get('password')
        level = request.args.get('level')

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("INSERT INTO user (username,password,level) VALUES ('{U}','{P}','{L}')"
                     .format(U=username, P=password, L=level))
        conn.commit()
        conn.close()

        return '', 200


class UserEdit(Resource):
    def get(self, user_id):
        abort_user(user_id)
        return get_user()[user_id]

    def post(self, user_id):
        abort_user(user_id)
        username = request.args.get('username')
        password = request.args.get('password')
        level    = request.args.get('level')
        id       = user_id[4:]

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("UPDATE user SET username = '{U}', password = '{P}', level = '{L}' WHERE id = {ID}"
                     .format(U=username, P=password, L=level, ID=id))

        conn.commit()
        conn.close()

        return get_user()[user_id]


class UserDelete(Resource):
    def post(self, user_id):
        abort_user(user_id)
        id = user_id[4:]
        conn = sqlite3.connect('mydatabase.db')
        conn.execute("DELETE FROM user WHERE id = {ID}".format(ID=id))

        conn.commit()
        conn.close()

        return get_user()


class PostAdd(Resource):
    def get(self):
        return get_post()

    def post(self):
        author      = request.args.get('username')
        title       = request.args.get('title')
        createdAt   = str(datetime.now())
        content     = request.args.get('content')

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("INSERT INTO post (author,title,createdAt,content) VALUES ('{A}','{T}','{C}','{CONTENT}')"
                     .format(A=author, T=title, C=createdAt, CONTENT=content))
        conn.commit()
        conn.close()

        return '', 200


api.add_resource(UserAdd,       '/user')
api.add_resource(UserEdit,      '/user/<user_id>')
api.add_resource(UserDelete,    '/deleteuser/<user_id>')
api.add_resource(PostAdd, '/post')

if __name__ == '__main__':
    app.run()
