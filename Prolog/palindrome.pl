palindrome(X) :- reverse(X, X).
palindrome_str(X) :- string_codes(X, Y), palindrome(Y).