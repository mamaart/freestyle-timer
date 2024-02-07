import requests, time

def new_session():
    r = requests.post(
        url = 'http://localhost:8080/new',
        json = {
            'session_title': 'first session',
            'athlete_one_name': 'martin',
            'athlete_two_name': 'bob',
        }
    )
    print(r)
    print(r.text)


def print_state():
    r = requests.get('http://localhost:8080/state')
    print(r)
    print(r.text)

def start(i):
    u = f'http://localhost:8080/start/{i}'
    print(u)
    r = requests.get(u)
    print(r)
    print(r.text)

def pause(i):
    u = f'http://localhost:8080/pause/{i}'
    print(u)
    r = requests.get(u)
    print(r)
    print(r.text)

if __name__ == "__main__":
    start(2)
