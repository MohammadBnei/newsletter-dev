FROM scratch

COPY newsletter-dev /app/newsletter

EXPOSE 8080

CMD ["/app/newsletter"]