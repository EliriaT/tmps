class Error:
    def __init__(self, name: str):
        self.name = name
    


errors = [
    Error('IllegalCharError'),
    Error('InvalidSyntaxError'),
    Error('RunTimeError')
]

def check_for_errors(errors: list):
    for error in errors:
        if error.name == 'IllegalCharError':
            print('You have an illegal character')

        elif error.name == 'InvalidSyntaxError':
            print('Please check your syntax')

        elif error.name == 'RunTimeError':
            print('You divided by 0')

