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


class UserAll(Resource):
    def get(self):
        return get_user()


class UserAdd(Resource):
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


class PostAll(Resource):
    def get(self):
        return get_post()


class PostAdd(Resource):
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


class PostEdit(Resource):
    def get(self, post_id):
        abort_post(post_id)
        return get_post()[post_id]

    def post(self, post_id):
        author      = request.args.get('username')
        title       = request.args.get('title')
        content     = request.args.get('content')
        id = post_id[4:]

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("UPDATE post SET author = '{A}', title = '{T}', content = '{CONTENT}' WHERE id = {ID}"
                     .format(A=author, T=title, CONTENT=content, ID=id))

        conn.commit()
        conn.close()

        return get_post()[post_id]


class PostDelete(Resource):
    def post(self, post_id):
        abort_post(post_id)
        id = post_id[4:]
        conn = sqlite3.connect('mydatabase.db')
        conn.execute("DELETE FROM post WHERE id = {ID}".format(ID=id))

        conn.commit()
        conn.close()

        return get_user()


class CategoryAll(Resource):
    def get(self):
        return get_category()


class CategoryAdd(Resource):
    def post(self):
        category = request.args.get('title')

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("INSERT INTO category (title) VALUES ('{C}')"
                     .format(C=category))
        conn.commit()
        conn.close()

        return '', 200


class CategoryEdit(Resource):
    def get(self, category_id):
        abort_category(category_id)
        return get_category()[category_id]

    def post(self, category_id):
        category = request.args.get('title')
        id       = category_id[8:]

        conn = sqlite3.connect('mydatabase.db')
        conn.execute("UPDATE category SET title = '{C}' WHERE id = {ID}"
                     .format(C=category, ID=id))

        conn.commit()
        conn.close()

        return get_category()[category_id]


class CategoryDelete(Resource):
    def post(self, category_id):
        abort_category(category_id)
        id = category_id[8:]
        conn = sqlite3.connect('mydatabase.db')
        conn.execute("DELETE FROM category WHERE id = {ID}".format(ID=id))

        conn.commit()
        conn.close()

        return get_category()


api.add_resource(UserAll, '/userlist')
api.add_resource(UserAdd, '/user')
api.add_resource(UserEdit, '/user/<user_id>')
api.add_resource(UserDelete, '/deleteuser/<user_id>')
api.add_resource(PostAll, '/postlist')
api.add_resource(PostAdd, '/post')
api.add_resource(PostEdit, '/post/<post_id>')
api.add_resource(PostDelete, '/deletepost/<post_id>')
api.add_resource(CategoryAll, '/categorylist')
api.add_resource(CategoryAdd, '/category')
api.add_resource(CategoryEdit, '/category/<category_id>')
api.add_resource(CategoryDelete, '/deletecategory/<category_id>')

if __name__ == '__main__':
    app.run()
