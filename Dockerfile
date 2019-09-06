FROM centos
WORKDIR /app
ENV PATH /app:$PATH
COPY sop /app/sop
COPY sop.db /app/sop.db
RUN chmod +x sop
CMD ["sop"]