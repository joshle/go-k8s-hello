# Scratch is a very small, minimalist docker image based on Debian.
FROM scratch

# Our app built as a binary executable.
ADD go-k8s-hello /

# Our app uses port 8080 by default.
EXPOSE 8080

# Run our app on startup.
CMD ["/go-k8s-hello"]