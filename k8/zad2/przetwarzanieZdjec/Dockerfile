FROM python:3.13.7

RUN apt-get update  \
    && apt-get upgrade -y \
    && pip install --upgrade pip \
    && apt-get install -y imagemagick unpaper parallel \
    && pip install page-dewarp numpy matplotlib \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN chmod +x procesuj.sh usun_numer_strony.py
ENTRYPOINT ["/bin/bash", "-c"]

# Komenda do wykonania przez ENTRYPOINT.
# To jest łańcuch znaków, który zostanie przekazany do /bin/bash -c
# To jest Twoja pożądana logika uruchamiana PRZY docker run.
CMD ["./procesuj.sh"]
# && python3 -m http.server 80