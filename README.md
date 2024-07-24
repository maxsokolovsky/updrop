# Updrop

This little web server allows you to copy and paste a piece of data
across the Internet or hosts on your local network.
It is useful when you are working with multiple computers and don't want to email things
to yourself or use cloud storage options for a quick share. You need to simply copy an arbitrary string from
one computer to another.

For security, updrop encrypts and decrypts data. It stores only one
string at a time in memory. You encrypt your string with a key.
When you decrypt using the key, the string is retrieved,
the store becomes empty, and the string cannot be retrieved again.

You use a key every time you want to share something across computers.
It must be of length 16, 24, or 32 bytes. You use the same key to decrypt.

You can host this anywhere - in the open or on your local network.

## Example

```
$ updrop

Listening on port :8000
```

You can configure the port:

```
$ updrop -port 6789
```

For help:

```
$ updrop -h
```
