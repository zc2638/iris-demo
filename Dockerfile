FROM centos
WORKDIR /app
ENV PATH /app:$PATH
COPY sop /app/sop
RUN chmod +x sop
CMD ["sop"]