# GOLANG API - ATENDIMENTO

## Descricao
Caso de estudo api golang. tentando se adaptar a clean arquiteture.
> Ferramentas:
> - Golang
> - RabbitMq
> - Sqlite

## Qual o problema ?
- Cliente: Quero que meu site tenha uma especie de atendimento e ouvidoria para dar suporte aos meus clientes.
- Site: Ja tem toda parte da visualizacao pronta, precisa apenas de endpoint's para a utilizacao dele.

## Qual a solucao ?
Criacao de uma api, para lidar **APENAS COM ATENDIMENTO**.

## Solicitacoes do cliente para o site da universidade.
1. Quero pode receber mensagens de reclamacoes e elogios dos meus alunos.
2. Desejo ter servicos especificos para meu atendimento. Como:
    - Solicitar trancamento de matricula
    - Solicitar a segunda chamada de provas
3. Quero que tenha passos para o aluno saber em que parte do atendimento ele ira estar.
   - Primeiro passo: ENVIADA SOLICITACAO
   - Segundo passo: ESPERA DA APROVACAO
   - Ultimo passo: SOLICITACAO FINALIZADA
4. Quem aprovara sera os coordenadores do curso no site de administracao da empresa.